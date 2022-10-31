function sendMail() {
  var xhr = new XMLHttpRequest();

  xhr.open('POST', '/api/sendMail');
  xhr.setRequestHeader('Content-Type', 'application/json');
  xhr.onload = function() {
    if (xhr.status === 200) {
      var res = JSON.parse(xhr.responseText);
      if (res.success == "true") {
        document.getElementById("sendButt").innerHTML = "Sent";
        document.getElementById("sendButt").style.background = "#7ca714";
      } else {
        document.getElementById("sendButt").innerHTML = "Error";
        document.getElementById("sendButt").style.background = "red";
      }
    }
  };

  xhr.send(JSON.stringify({
        email: document.getElementById("emailInput").value,
        phoneNumber: document.getElementById("phoneNumber").value,
        message: document.getElementById("messageArea").value,
  }));
}
