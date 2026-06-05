const itemsPerPage = 12;
let currentPage = 1;
let currentOwner = null;

document.addEventListener("DOMContentLoaded", () => {
  const urlParams = new URLSearchParams(window.location.search);
  currentOwner = urlParams.get("owner");

  // 初始化页面
  renderPage();

  // 添加分页按钮事件监听
  document.getElementById('prevBtn').addEventListener('click', function () {
    if (currentPage > 1) {
      currentPage--;
      renderPage();
    }
  });

  document.getElementById('nextBtn').addEventListener('click', function () {
    currentPage++;
    renderPage();
  });
});

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // 清空当前内容
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listVolume(currentOwner, startIndex, itemsPerPage)
    .then((data) => {
      if (data.length === 0 && currentPage > 1) {
        // 如果当前页没有数据且不是第一页，返回上一页
        currentPage--;
        renderPage();
        return;
      }
      data.forEach((meta) => {
        const card = createCard(meta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching volume:", error));

  // 更新页码显示
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = meta.Piece;
  card.appendChild(h2);

  const cardContent = document.createElement("div");
  cardContent.className = "card-content";
  card.appendChild(cardContent);

  const p100 = document.createElement("p");
  p100.setAttribute('data-label', 'ChainType:');
  p100.textContent = meta.ChainType;
  cardContent.appendChild(p100);

  const p1 = document.createElement("p");
  p1.setAttribute('data-label', 'Owner:');
  p1.textContent = meta.Owner;
  cardContent.appendChild(p1);

  const p2 = document.createElement("p");
  p2.setAttribute('data-label', 'Volume:');
  p2.textContent = meta.File;
  cardContent.appendChild(p2);

  const p3 = document.createElement("p");
  p3.setAttribute('data-label', 'Creation:');
  p3.textContent = meta.CreatedAt;
  cardContent.appendChild(p3);

  if (meta.TxHash) {
    let rurl = "https://sepolia-optimism.etherscan.io/tx/"
    if (meta.ChainType == "bnb-testnet" || meta.ChainType == "bnb-testnet-v2" || meta.ChainType == "bnb-testnet-dao") {
      rurl = "https://testnet.bscscan.com/tx/"
    } else if (meta.ChainType == "opbnb-testnet") {
      rurl = "https://opbnb-testnet.bscscan.com/tx/"
    }
    const p5 = document.createElement("p");
    p5.setAttribute('data-label', 'TxHash:');
    const link = document.createElement('a');
    link.href = rurl + meta.TxHash;
    link.target = '_blank';
    link.textContent = meta.TxHash;
    p5.appendChild(link);
    cardContent.appendChild(p5);
  }

  return card;
}

function listVolume(eaddr, offset = 0, length = itemsPerPage) {
  if (eaddr) {
    return fetch(`/api/listVolume?owner=${eaddr}&offset=${offset}&length=${length}`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching volume:", error))
  } else {
    return fetch(`/api/listVolume?offset=${offset}&length=${length}`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching volume:", error))
  }
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getVolume?owner=${searchInput}`)
    .then((response) => response.json())
    .then((data) => {
      displayResults(data);
    })
    .catch((error) => console.error("Error:", error));
}

function displayResults(data) {
  const resultsElement = document.getElementById("cardContainer");
  resultsElement.innerHTML = ""; // 清空先前的结果

  if (data.length === 0) {
    resultsElement.innerText = "No results found.";
  } else {
    data.forEach((meta) => {
      const card = createCard(meta);
      resultsElement.appendChild(card);
    });
  }

  // 重置分页状态
  currentPage = 1;
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}