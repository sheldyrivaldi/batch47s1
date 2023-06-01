// Menampilkan nama file ketika sudah memilih image
let inputElement = document.getElementById("upload-image");
let fileNameElement = document.getElementById("file-name");

inputElement.addEventListener("change", () => {
  let files = inputElement.files;
  if (files.length > 0) {
    let fileName = files[0].name;
    fileNameElement.textContent = fileName;
  }
});

//Mensubmit Data Project
let dataProject = [];
let dataTechnology = [];

function submitData(event) {
  event.preventDefault();

  // Deklarasi variable dari input
  let title = document.getElementById("project-name").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("description").value;
  let nodejs = document.getElementById("nodejs");
  let nextjs = document.getElementById("nextjs");
  let reactjs = document.getElementById("reactjs");
  let typescript = document.getElementById("typescript");
  let image = document.getElementById("upload-image").files;

  // Mencari durasi project
  let mulai = new Date(startDate);
  let akhir = new Date(endDate);
  let selisih = akhir.getTime() - mulai.getTime();
  let days = selisih / (1000 * 60 * 60 * 24);
  let weeks = Math.floor(days / 7);
  let months = Math.floor(weeks / 4);
  let years = Math.floor(months / 12);
  let durasi = "";

  if (days > 0) {
    durasi = days + " hari";
  }
  if (weeks > 0) {
    durasi = weeks + " minggu";
  }
  if (months > 0) {
    durasi = months + " bulan";
  }
  if (years > 0) {
    durasi = years + " tahun";
  }

  // Cek Checkbox
  if (nodejs.checked) {
    dataTechnology.push(nodejs.value);
  }
  if (nextjs.checked) {
    dataTechnology.push(nextjs.value);
  }
  if (reactjs.checked) {
    dataTechnology.push(reactjs.value);
  }
  if (typescript.checked) {
    dataTechnology.push(typescript.value);
  }

  // Membuat url image
  let imageUrl = URL.createObjectURL(image[0]);

  // Membuat object data project
  let data = {
    title,
    startDate,
    endDate,
    duration: durasi,
    description,
    technology: dataTechnology,
    imageUrl,
  };
  //   dataProject.push(data);
  console.log(data);
  dataTechnology = [];
  dataProject = [];
}
