package hub

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	contract "github.com/unibaseio/da-sdk-go/contract/v2"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/gin-gonic/gin"
)

func (s *Server) addUpload(g *gin.RouterGroup) {
	g.Group("/").POST("/uploadData", s.uploadData)
	g.Group("/").POST("/upload", s.upload)
}

func (s *Server) uploadData(c *gin.Context) {
	addr := c.PostForm("owner")
	if addr == "" {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("owner is required")))
		return
	}
	bucket := c.PostForm("bucket")
	if bucket == "" {
		bucket = addr
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if file == nil {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("file is nil")))
		return
	}

	fr, err := file.Open()
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if file.Size == 0 {
		c.JSON(599, lerror.ToAPIError("hub", fmt.Errorf("empty file")))
		return
	}
	mm, err := s.logFSWrite(addr, bucket, file.Filename, fr)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, mm)
}

func (s *Server) upload(c *gin.Context) {
	var mjson types.MemeStruct

	err := c.ShouldBindJSON(&mjson)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	if mjson.Bucket == "" {
		var meta map[string]interface{}
		err = json.Unmarshal([]byte(mjson.Message), &meta)
		if err != nil {
			c.JSON(599, lerror.ToAPIError("hub", err))
			return
		}
		bucketName, ok := meta["name"].(string)
		if ok {
			mjson.Bucket = bucketName
		} else {
			mjson.Bucket = mjson.Owner
		}
	}

	var buf bytes.Buffer
	buf.WriteString(mjson.Message)

	mm, err := s.logFSWrite(mjson.Owner, mjson.Bucket, mjson.ID, &buf)
	if err != nil {
		c.JSON(599, lerror.ToAPIError("hub", err))
		return
	}

	c.JSON(http.StatusOK, mm)
}

func (s *Server) logFSWrite(addr string, bucket string, key string, r io.Reader) (types.MemeMeta, error) {
	var err error
	if addr == "" {
		addr = s.local.String()
	}

	if bucket == "" {
		bucket = addr
	}

	err = s.addBucket(addr, bucket)
	if err != nil {
		return types.MemeMeta{}, err
	}

	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		fspath := filepath.Join(s.rp.Path(), LOGFS)
		fs, err = logfs.New(s.rp.MetaStore(), fspath, s.local.String(), addr)
		if err != nil {
			s.Unlock()
			return types.MemeMeta{}, err
		}
		s.lfs[addr] = fs
		s.addAccount(addr)
		logger.Infof("start new log inst: %s %d", addr, s.fscnt)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, addr)
		has, err := s.rp.MetaStore().Has(dsKey)
		if err != nil || !has {
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, 0)
			s.rp.MetaStore().Put(dsKey, buf)

			dsKey = types.NewKey(types.DsLogFS, LOGINST, s.fscnt)
			s.rp.MetaStore().Put(dsKey, []byte(addr))

			dsKey = types.NewKey(types.DsLogFS, LOGINST)
			s.fscnt++
			binary.BigEndian.PutUint32(buf, s.fscnt)
			s.rp.MetaStore().Put(dsKey, buf)
		}
	}
	s.Unlock()
	rbytes, err := io.ReadAll(r)
	if err != nil {
		return types.MemeMeta{}, err
	}

	err = fs.Put([]byte(key), rbytes)
	if err != nil {
		return types.MemeMeta{}, err
	}

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return types.MemeMeta{}, err
	}

	s.addNeedle(addr, bucket, key, lm.Index, lm.Start, lm.Size)

	mm := types.MemeMeta{
		File:  fmt.Sprintf("%s/%d.log", addr, lm.Index),
		Start: lm.Start,
		Size:  lm.Size,
	}

	return mm, nil
}

func (s *Server) logFSRead(addr string, key string, w io.Writer) (int64, error) {
	var err error
	if addr == "" {
		ns, err := s.getNeedleByName(key)
		if err != nil || len(ns) == 0 {
			return 0, fmt.Errorf("no such needle %s", key)
		} else {
			addr = ns[0].Owner
		}
	}

	s.Lock()
	fs, ok := s.lfs[addr]
	if !ok {
		dsKey := types.NewKey(types.DsLogFS, LOGINST, addr)
		has, err := s.rp.MetaStore().Has(dsKey)
		if err == nil && has {
			fspath := filepath.Join(s.rp.Path(), LOGFS)
			fs, err = logfs.New(s.rp.MetaStore(), fspath, s.local.String(), addr)
			if err != nil {
				s.Unlock()
				return 0, err
			}
			s.lfs[addr] = fs
			logger.Infof("start exist log inst: %s %d", addr, s.fscnt)
		} else {
			s.Unlock()
			return 0, fmt.Errorf("no such owner: %s", addr)
		}
	}
	s.Unlock()

	lm, err := fs.GetMeta([]byte(key))
	if err != nil {
		return 0, err
	}

	wbytes, err := fs.GetData(lm)
	if err != nil {
		// todo: get from remote
		return 0, err
	}
	n, err := w.Write(wbytes)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}

func (s *Server) load() error {
	fspath := filepath.Join(s.rp.Path(), LOGFS)
	fs, err := logfs.New(s.rp.MetaStore(), fspath, s.local.String(), s.local.String())
	if err != nil {
		return err
	}
	s.lfs[s.local.String()] = fs

	dsKey := types.NewKey(types.DsLogFS, LOGINST)
	val, err := s.rp.MetaStore().Get(dsKey)
	if err == nil && len(val) == 4 {
		s.fscnt = binary.BigEndian.Uint32(val)
	}

	if s.fscnt == 0 {
		s.fscnt = 1
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, s.fscnt)
		s.rp.MetaStore().Put(dsKey, buf)

		dsKey := types.NewKey(types.DsLogFS, LOGINST, 0)
		s.rp.MetaStore().Put(dsKey, []byte(s.local.String()))
	}
	/*
		for i := uint32(0); i < s.fscnt; i++ {
			dsKey := types.NewKey(types.DsLogFS, LOGINST, i)
			val, err := s.rp.MetaStore().Get(dsKey)
			if err != nil {
				break
			}

			s.addAccount(string(val))
		}
	*/
	logger.Infof("load log inst: %d", s.fscnt)
	return nil
}

func (s *Server) uploadTo() {
	sk := s.rp.Key().Export().PrivateKey
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		panic(err)
	}

	policy := types.Policy{
		N: 6,
		K: 4,
	}

	cm, err := contract.NewContractManage(sk, s.rp.Repo().Config().Chain.Type)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Minute)
		logger.Info("check uploaded info")
		err := cm.CheckBalance(au.Addr)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}

		logger.Info("check log inst: ", s.fscnt)

		for i := uint32(0); i < s.fscnt; i++ {
			dsKey := types.NewKey(types.DsLogFS, LOGINST, i)
			val, err := s.rp.MetaStore().Get(dsKey)
			if err != nil {
				break
			}

			key := string(val)

			logger.Debugf("check: %s %d", key, i)
			dsKey = types.NewKey(types.DsLogFS, key)
			val, err = s.rp.MetaStore().Get(dsKey)
			if err != nil || len(val) != 8 {
				continue
			}
			curIndex := binary.BigEndian.Uint64(val)

			next := logfs.GetIndex(s.local.String(), key)
			dsKey = types.NewKey(types.DsLogFS, LOGINST, key)
			val, err = s.rp.MetaStore().Get(dsKey)
			if err == nil && len(val) == 8 {
				next = binary.BigEndian.Uint64(val)
			}

			logger.Debugf("check: %s %d %d", key, next, curIndex)
			if next >= curIndex {
				continue
			}

			for i := next; i < curIndex; i++ {
				fname := fmt.Sprintf("%s/%d.vol", key, i)
				fp := filepath.Join(s.rp.Path(), LOGFS, key, fmt.Sprintf("%d.vol", i))

				fr, err := sdk.GetFileReceipt(sdk.ServerURL, au, fname)
				if err == nil {
					logger.Infof("%s/%d.vol is already uploaded, check its piece onchain", key, i)
					if fr.ChainType != s.rp.Repo().Config().Chain.Type {
						buf := make([]byte, 8)
						binary.BigEndian.PutUint64(buf, i+1)
						s.rp.MetaStore().Put(dsKey, buf)
						logger.Warnf("new chain type detected, ignore previous one")
						continue
					}
					er, err := sdk.ListEdge(sdk.ServerURL, au, types.StreamType)
					if err != nil {
						break
					}
					suc := 0
					for _, pn := range fr.Pieces {
						for _, st := range er.Edges {
							pr, err := sdk.GetPieceReceipt(st.ExposeURL, au, pn)
							if err == nil {
								if pr.Serial > 0 {
									suc++
								} else {
									txn, err := cm.AddPiece(pr.PieceCore)
									if err == nil {
										s.addVolume(key, i, pr.Name, txn)
										suc++
									}
								}
							}
						}
					}
					if suc == len(fr.Pieces) {
						buf := make([]byte, 8)
						binary.BigEndian.PutUint64(buf, i+1)
						s.rp.MetaStore().Put(dsKey, buf)
						continue
					}
					continue
				}
				// upload to stream and submit to gateway
				res, streamer, err := sdk.Upload(sdk.ServerURL, au, policy, fp, fname)
				if err != nil {
					if strings.Contains(err.Error(), "already has piece") {
						logger.Warnf("piece of file %s found on server, skip", fp)
						continue
					}
					break
				}
				pcs, err := sdk.CheckFileFull(res, streamer, fp)
				if err != nil {
					break
				}
				log.Printf("upload %s to %s, sha256: %s\n", fp, streamer, res.Hash)
				log.Printf("submit %s to chain\n", res.Name)
				// submit meta to chain
				var terr error
				for _, pc := range pcs {
					txn, err := cm.AddPiece(pc)
					if err != nil {
						terr = err
						break
					}
					s.addVolume(key, i, pc.Name, txn)
				}
				if terr != nil {
					break
				}
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, i+1)
				s.rp.MetaStore().Put(dsKey, buf)
			}
		}
	}
}
