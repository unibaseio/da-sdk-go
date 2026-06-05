package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate/entities/models"
)

const classname = "Vector_index_b336533e_f79e_4e58_a7e1_546492f7b0be_Node"
const tenantname = "7249ca19-3c38-4ca0-b1bf-c36c46f1c01b"
const newtenant = "7249ca19-3c38-4ca0-b1bf-c36c46f1c01b"
const weaviatehost = "18.141.185.111:8080"
const hubhost = "http://54.151.130.2:8080"

func GetSchema() error {
	cfg := weaviate.Config{
		Host:   weaviatehost,
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return err
	}

	ctx := context.TODO()

	class := &models.Class{
		Class: classname,
		//MultiTenancyConfig: &models.MultiTenancyConfig{
		//	Enabled:            true,
		//	AutoTenantCreation: true,
		//},
	}

	cl, err := client.Schema().ClassGetter().WithClassName(classname).Do(ctx)
	if err == nil {
		b, err := json.MarshalIndent(cl, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	} else {
		err = client.Schema().ClassCreator().WithClass(class).Do(ctx)
		if err != nil {
			return err
		}
	}

	tenants, err := client.Schema().TenantsGetter().
		WithClassName(classname).
		Do(ctx)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(tenants, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func Migrate(host, cn, old, new string) error {
	cfg := weaviate.Config{
		Host:   host,
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return err
	}

	ctx := context.TODO()

	mo, err := client.Data().ObjectsGetter().WithClassName(cn).WithTenant(old).WithVector().Do(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < len(mo); i++ {
		w, err := client.Data().Creator().
			WithClassName(cn).
			WithProperties(mo[i].Properties).
			WithTenant(new).
			WithVectors(mo[i].Vectors).
			WithVector(mo[i].Vector).
			WithID(mo[i].ID.String()).
			Do(ctx)
		if err != nil {
			return err
		}
		fmt.Println(mo[i].ID, " migrate to:", w.Object.ID, w.Object.Tenant)
	}

	return nil
}

func TestGet(t *testing.T) {
	err := Migrate(weaviatehost, classname, tenantname, newtenant)
	t.Fatal(err)
}

func TestCreate(t *testing.T) {
	cfg := weaviate.Config{
		Host:       weaviatehost,
		Scheme:     "http",
		AuthConfig: auth.ApiKey{Value: "WVF5YThaHlkYwhGUSmCRgsX3tD5ngdN8pkih"},
		//@Headers:    nil,
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	w, err := client.Data().Creator().
		WithClassName(classname).
		WithProperties(map[string]interface{}{
			"question":    "This vector DB is OSS and supports automatic property type inference on import",
			"answer":      "Weaviate1", // schema properties can be omitted
			"newProperty": 12356789,    // will be automatically added as a number property
		}).
		WithTenant(tenantname).
		Do(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	// the returned value is a wrapped object
	b, err := json.MarshalIndent(w.Object, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))

	time.Sleep(5 * time.Second)

	resb, err := sdk.DownloadHubData(hubhost, w.Object.Tenant, w.Object.ID.String())
	if err != nil {
		t.Fatal(err)
	}

	res := new(models.Object)
	err = res.UnmarshalBinary(resb)
	if err != nil {
		t.Fatal(err)
	}

	b, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
