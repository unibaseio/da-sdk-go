const urlParams = new URLSearchParams(window.location.search);
let eaddr = ""
const eaddrparam = urlParams.get("owner");
if (eaddrparam) {
  eaddr = eaddrparam;
}
let bucket = ""
const bucketparam = urlParams.get("bucket");
if (bucketparam) {
  bucket = bucketparam;
}
let conversation = ""
const conversationparam = urlParams.get("conversation");
if (conversationparam) {
  conversation = conversationparam;
}

const itemsPerPage = 12;
let currentPage = 1;

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // Clear the current content
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listNeedle(eaddr, startIndex, itemsPerPage)
    .then((data) => {
      data.forEach((fileMeta) => {
        const card = createCard(fileMeta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching needle:", error));

  // Update the page number
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}

// Handle "Previous" button click
document.getElementById('prevBtn').addEventListener('click', function () {
  if (currentPage > 1) {
    currentPage--;
    renderPage();
  }
});

// Handle "Next" button click
document.getElementById('nextBtn').addEventListener('click', function () {
  currentPage++;
  renderPage();
});

// Initial page render
renderPage();

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = meta.Name;
  card.appendChild(h2);

  const cardContent = document.createElement("div");
  cardContent.className = "card-content";
  card.appendChild(cardContent);

  const p5 = document.createElement("p");
  p5.setAttribute('data-label', 'Owner:');
  p5.textContent = meta.Owner;
  cardContent.appendChild(p5);

  const p7 = document.createElement("p");
  p7.setAttribute('data-label', 'Agent:');
  p7.textContent = meta.Bucket;
  cardContent.appendChild(p7);

  const p0 = document.createElement("p");
  p0.setAttribute('data-label', 'Volume:');
  p0.textContent = meta.File;
  cardContent.appendChild(p0);

  const p1 = document.createElement("p");
  p1.setAttribute('data-label', 'Start:');
  p1.textContent = meta.Start;
  cardContent.appendChild(p1);

  const p2 = document.createElement("p");
  p2.setAttribute('data-label', 'Size:');
  p2.textContent = meta.Size + ' bytes';
  cardContent.appendChild(p2);

  const p21 = document.createElement("p");
  p21.setAttribute('data-label', 'Creation:');
  p21.textContent = meta.CreatedAt;
  cardContent.appendChild(p21);

  if (meta.TxHash) {
    const p3 = document.createElement("p");
    p3.setAttribute('data-label', 'Piece:');
    p3.textContent = meta.Piece;
    cardContent.appendChild(p3);

    let rurl = "https://sepolia-optimism.etherscan.io/tx/"
    if (meta.ChainType == "bnb-testnet" || meta.ChainType == "bnb-testnet-v2" || meta.ChainType == "bnb-testnet-dao") {
      rurl = "https://testnet.bscscan.com/tx/"
    } else if (meta.ChainType == "opbnb-testnet") {
      rurl = "https://opbnb-testnet.bscscan.com/tx/"
    }

    const p4 = document.createElement("p");
    p4.setAttribute('data-label', 'TxHash:');
    const link = document.createElement('a');
    link.href = rurl + meta.TxHash;
    link.target = '_blank';
    link.textContent = meta.TxHash;
    p4.appendChild(link);
    cardContent.appendChild(p4);
  }

  const p6 = document.createElement("p");
  p6.setAttribute('data-label', 'Content:');
  const contentLink = document.createElement('a');
  contentLink.href = 'content.html?owner=' + meta.Owner + '&name=' + meta.Name;
  contentLink.target = '_blank';
  contentLink.textContent = '[click to show]';
  p6.appendChild(contentLink);
  cardContent.appendChild(p6);

  return card;
}

function listNeedle(eaddr, offset, length) {
  return fetch(`/api/listNeedle?owner=${eaddr}&bucket=${bucket}&conversation=${conversation}&offset=${offset}&length=${length}`)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("Error fetching needles:", error))
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getNeedle?owner=${eaddr}&bucket=${bucket}&conversation=${conversation}&name=${searchInput}`)
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
      cardContainer.appendChild(card);
    });
  }
}


