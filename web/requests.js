baseUrl = "http://localhost:10000"

function pollForGame() {
	// continuously poll backend for updates every X ms
	setInterval( function() {
		updateGame()
	}, 500) // in ms
}

function newGame(boardSize) {
	var urlExtension = "/new_game?board_size=" + boardSize.toString();
	baseRequestDraw(urlExtension);
}

function getGame() {
	var urlExtension = "/games?game_id=0";
	baseRequestDraw(urlExtension);
}

function updateGame() {
	var urlExtension = "/games?game_id=0";
	baseRequest(urlExtension);
}

function resetGame() {
	var urlExtension = "/reset_game?game_id=0";
	baseRequest(urlExtension);
}

function placePiece(cell) {
	cell.classList.remove("hovering");
	var row = cell.id.split("-")[0];
	var col = cell.id.split("-")[1];
	var urlExtension = "/place_piece?game_id=0&row=" + row + "&col=" + col;
	baseRequest(urlExtension);
}

function baseRequest(urlExtension) {
	var xmlhttp = new XMLHttpRequest();
	var url = baseUrl + urlExtension;

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			game = JSON.parse(this.responseText);
			updateBoard(game["Board"]);
		}
	};

	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}

function baseRequestDraw(urlExtension) {
	var xmlhttp = new XMLHttpRequest();
	var url = baseUrl + urlExtension;

	xmlhttp.onreadystatechange = function() {
		if (this.status === 200 && this.readyState === XMLHttpRequest.DONE) {
			console.log(this.responseText);
			game = JSON.parse(this.responseText);
			drawBoard(game["Board"]);
		}
	};

	xmlhttp.open("GET", url, true);
	xmlhttp.send();
}
