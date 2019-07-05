game = {}
baseUrl = "http://localhost:10000"

function hello() {
    let myElem = document.getElementById("my-elem");
    myElem.innerHTML = "Hello, World!";
	// pollForGame() //(ngrok rate limit reached real quickly lol)
}

function pollForGame() {
	var xmlhttp = new XMLHttpRequest();
	var url = baseUrl + "/games?game_id=0";

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			game = JSON.parse(this.responseText);
			drawBoard(game["Board"]);
		}
	};
	
	xmlhttp.open("GET", url, true);
	xmlhttp.send();
	setInterval( function() {
		xmlhttp.open("GET", url, true);
		xmlhttp.send();
		// continuously poll backend for updates every X ms
	}, 500)
}

function getGame() {
	var xmlhttp = new XMLHttpRequest();
	var url = baseUrl + "/games?game_id=0";

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			game = JSON.parse(this.responseText);
			drawBoard(game["Board"]);
		}
	};
	
	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}

function resetGame() {
	var xmlhttp = new XMLHttpRequest();
	var url = baseUrl + "/reset_game?game_id=0";

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			game = JSON.parse(this.responseText);
			drawBoard(game["Board"]);
		}
	};
	
	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}

function placePiece(cell) {
	console.log("placing piece...")
	var xmlhttp = new XMLHttpRequest();
	var row = cell.id.split("-")[0];
	var col = cell.id.split("-")[1];
	var url = baseUrl + "/place_piece?game_id=0&row=" + row + "&col=" + col;

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			console.log(this.responseText);
			console.log(game);
			game = JSON.parse(this.responseText);
			drawBoard(game["Board"]);
		}
	};

	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}

function getCellSize(boardSize) {
	var boardSizeInPx = 500;
	var cellSizeInPx = boardSizeInPx / boardSize;
	return Math.round(cellSizeInPx.toString()) + "px"
}

function hoverPiece(cell) {
	var cellImg = cell.firstChild;
	if (game["CurrentPlayer"] === "dark") {
		cellImg.src = "assets/black-circle.jpeg";
	} else if (game["CurrentPlayer"] === "light") {
		cellImg.src = "assets/white-circle.png";
	}
}

function unHoverPiece(cell) {
	var cellImg = cell.firstChild;
	cellImg.src = "assets/blank-square.png";
}

function drawBoard(board) {
	var boardSize = board.length;
	var tableDiv = document.getElementById("board");
	tableDiv.innerHTML = "";
	var tbl = document.createElement("table");
	for (var r = 0; r < boardSize; r++) {
		var row = document.createElement("tr");

		for (var c = 0; c < boardSize; c++) {
			var cell = document.createElement("td");
			var cellSize = getCellSize(boardSize);
			cell.style.width = cellSize;
			cell.style.height = cellSize;
			cell.id = r.toString() + "-" + c.toString();
			var cellImg = document.createElement("img");
			if (board[r][c] === 1) {
				cellImg.src = "assets/black-circle.jpeg";
			} else if (board[r][c] === 2) {
				cellImg.src = "assets/white-circle.png";
			} else {
				cellImg.src = "assets/blank-square.png";
				cell.onmouseover = function() { hoverPiece(this) }
				cell.onmouseout = function() { unHoverPiece(this) }
				cell.onclick = function() { placePiece(this) }
			}
			cell.appendChild(cellImg);
			row.appendChild(cell);
		}
		tbl.appendChild(row);
	}
	tableDiv.appendChild(tbl);
}
