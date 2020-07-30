var $table = $('#table');
const serverDomain = 'http://localhost:8080/'


let data = JSON.parse(http(serverDomain + 'instruments', "GET"))


function http(theUrl, method, data) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open(method, theUrl, false);
    if (method === "POST") {
        xmlHttp.setRequestHeader('Content-Type', 'application/json');
    }
    xmlHttp.send(JSON.stringify(data));
    return xmlHttp.responseText;
}

$(function () {
    $('#table').bootstrapTable({
        data: data
    });
});

function deleteInstrument() {
    let instID = document.getElementById("delInst");
    http(serverDomain + 'deleteinstrument?id=' + instID.value, "DELETE")
    location.reload();
}

function addInstrument() {
    let instID = document.getElementById("instID").value;
    let instName = document.getElementById("instName").value;
    let instType = document.getElementById("instType").value;
    let symbol = document.getElementById("symbol").value;
    let data = {
        instrumentId: parseInt(instID),
        name: instName,
        symbol: symbol,
        instrumentType: instType
    };

    http(serverDomain + 'addinstrument', "POST", data)
    location.reload();
}