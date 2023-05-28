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
    let image = document.getElementById("add-project-upload-image").files;

    // Mencari durasi project
    let mulai = new Date(startDate);
    let akhir = new Date(endDate);
    let selisih = akhir.getTime() - mulai.getTime();
    let days = selisih / (1000 * 60 * 60 * 24)
    let weeks = Math.floor(days / 7)
    let months = Math.floor(weeks / 4)
    let years = Math.floor(months / 12)
    let durasi = ""

    if(days > 0){
        durasi = days + " hari"
    }
    if (weeks > 0){
        durasi = weeks + " minggu"
    }
    if (months > 0){
        durasi = months + " bulan"
    }
    if (years > 0){
        durasi = years + " tahun"
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
        durasi,
        description,
        dataTechnology,
        imageUrl,
    };
    dataProject.push(data);
    renderProject();
    dataTechnology = []
}

function renderProject() {
    document.getElementById("project-list").innerHTML = "";

    for (let i = 0; i < dataProject.length; i++) {
        let technologyImages = "";

        for (let j = 0; j < dataProject[i].dataTechnology.length; j++) {
            technologyImages += `<img src="assets/images/${dataProject[i].dataTechnology[j]}.png" alt="${dataProject[i].dataTechnology[j]}">`;
        }

        document.getElementById("project-list").innerHTML += 
            ` <div id="project-items" class="project-items">
            <div class="project-items-container">
                <div class="project-list-image">
                    <img src="${dataProject[i].imageUrl}" alt="project-list">
                </div>
                <div class="project-list-title">
                    <p class="list-title"><a target="_blank" href="project-detail.html">${dataProject[i].title}s</a></p>
                    <p class="list-duration">durasi : ${dataProject[i].durasi}</p>
                </div>
                <div class="description">
                    <p class="list-description">${dataProject[i].description}</p>
                </div>
                <div class="technology">
                    ${technologyImages}
                </div>
                <div class="project-list-button">
                    <button class="edit" type="button">edit</button>
                    <button class="delete" type="button">delete</button>
                </div>
            </div>
        </div>`;
    }
}

let humbergerIsOpen = false;

function navbarSwitch(){
   const humbergerNavbarList =  document.getElementById("navbar-humberger-list");
   if (!humbergerIsOpen){
        humbergerNavbarList.style.display = "block";
        humbergerIsOpen = true;
   } else {
        humbergerNavbarList.style.display = "none";
        humbergerIsOpen = false;
   }
}
