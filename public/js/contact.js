let subjectItem = "";
let dropdownItems = ["Frontend Developer", "Backend Developer", "Fullstack Developer"];

function dropdownClick(item) {
  subjectItem = dropdownItems[item];
}

function submitData() {
  //Deklarasi variabel
  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;
  let phone = document.getElementById("phone").value;
  let subject = subjectItem;
  let message = document.getElementById("message").value;

  //Validasi input data
  if (name == "") {
    return alert("Nama tidak boleh kosong!");
  }
  if (email == "") {
    return alert("Email tidak boleh kosong!");
  }
  if (phone == "") {
    return alert("Nomor HP tidak boleh kosong!");
  }
  if (subject == "") {
    return alert("Subject tidak boleh kosong!");
  }

  //Link "mailto"
  let recivedEmail = "sheldyrivaldi@gmail.com";
  let enter = "\n\n";
  let a = document.createElement("a");
  let mailToLink = `mailto:${recivedEmail}?subject=${encodeURIComponent(subject)}&body=Halo! Nama saya ${encodeURIComponent(name)}.${encodeURIComponent(enter)}${encodeURIComponent(message)}.${encodeURIComponent(
    enter
  )} Kamu bisa mengubungi saya pada nomor ${encodeURIComponent(phone)}.${encodeURIComponent(enter)} Terimakasih.`;
  a.href = mailToLink;

  a.click();
}
