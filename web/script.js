game = {}

function hello() {
    let myElem = document.getElementById("my-elem");
    myElem.innerHTML = "Go in Go!";
	getGame();
	pollForGame()
}

function calculateCellSize(boardSize) {
	// calculates how wide (in pixels) the squares should be based on the board size
	var boardSizeInPx = 500;
	var cellSizeInPx = boardSizeInPx / boardSize;
	return Math.round(cellSizeInPx.toString()) + "px"
}

function hoverPiece(cell) {
	cell.classList.add("hovering");
	var cellImg = cell.firstChild;
	if (game["CurrentPlayer"] === "dark") {
		cellImg.src = "assets/black-circle.jpeg";
	} else if (game["CurrentPlayer"] === "light") {
		cellImg.src = "assets/white-circle.png";
	}
}

function unHoverPiece(cell) {
	cell.classList.remove("hovering");
	var cellImg = cell.firstChild;
	cellImg.src = "assets/blank-square.png";
}

function drawBoard(board) {
	var boardSize = board.length;
	var tableDiv = document.getElementById("board");
	var tbl = document.createElement("table");

	for (var i = 0; i < boardSize; i++) {
		var row = document.createElement("tr");

		for (var j = 0; j < boardSize; j++) {
			var cell = document.createElement("td");
			var cellSize = calculateCellSize(boardSize);
			cell.style.width = cellSize;
			cell.style.height = cellSize;

			// store the index of each square in its id, e.g. "0-0" is the id of the first square
			cell.id = i.toString() + "-" + j.toString();

			var cellImg = document.createElement("img");
			if (board[i][j] === 1) {
				cellImg.src = "assets/black-circle.jpeg";
			} else if (board[i][j] === 2) {
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
	tableDiv.innerHTML = ""; // removes the previous board
	tableDiv.appendChild(tbl);
}

function updateBoard(board) {
	var boardSize = board.length;

	for (var i = 0; i < boardSize; i++) {
		for (var j = 0; j < boardSize; j++) {
			var cellId = i.toString() + "-" + j.toString();
			var cell = document.getElementById(cellId);

			var cellImg = cell.firstChild;
			if (board[i][j] === 1) {
				cellImg.src = "assets/black-circle.jpeg";
				removeMouseEvents(cell);
			} else if (board[i][j] === 2) {
				cellImg.src = "assets/white-circle.png";
				removeMouseEvents(cell);
			} else {
				if (cell.classList.contains("hovering")) {
					return
				}
				cellImg.src = "assets/blank-square.png";
				addMouseEvents(cell);
			}
		}
	}
}

function addMouseEvents(cell) {
	cell.onmouseover = function() { hoverPiece(this) }
	cell.onmouseout = function() { unHoverPiece(this) }
	cell.onclick = function() { placePiece(this) }
}

function removeMouseEvents(cell) {
	cell.onmouseover = function() {}
	cell.onmouseout = function() {}
	cell.onclick = function() {}
}
