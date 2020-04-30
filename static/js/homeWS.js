window.onload = function () {
    var room_name;
    var get_room_name = function (){
        var url = document.location.toString().split("/room/").slice(-1);
        room_name = url;
        return room_name
    };

    var conn = new WebSocket("ws://" + document.location.host + "/api/v1/room/"  + get_room_name() + "/ws"), input_msg = document.getElementById("input_msg");



    var appendMSG = function(str) {
        var
            msg_board = document.getElementById("msg_board"),
            last_msg = document.getElementById("last_msg"),
            new_msg = document.createElement("span");

        new_msg.innerHTML = str;
        new_msg.className = "single_msg";
        msg_board.insertBefore(document.createElement("br"),last_msg);
        msg_board.insertBefore(new_msg, last_msg);

    };
    conn.onmessage = function (event) {
        var str = event.data;
        appendMSG(str)
    };
    conn.onclose = function (event) {
        var str = "<h3>Connection has closed.</h3>";
        appendMSG(str)
    };

    document.getElementById("send_box").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!input_msg.value) {
            return false;
        }
        conn.send(input_msg.value);
        input_msg.value = "";
        return false;
    };
};