const Data = [
    {
        author: "Surya Elidanto",
        rating: 5,
        opinion: "Keren banget jasanya!",
        image: "assets/images/testimonial-3.jpg"
    },
    {
        author: "Surya Elz",
        rating: 4,
        opinion: "Keren lah pokonya!",
        image: "assets/images/testimonial-2.jpg"
    },
    {
        author: "Surya Gans",
        rating: 4,
        opinion: "The best pelayanannya!",
        image: "assets/images/testimonial-1.jpg"
    },
    {
        author: "Suryaaaa",
        rating: 3,
        opinion: "Oke lah!",
        image: "assets/images/testimonial-3.jpg"
    },
    {
        author: "Suryeah",
        rating: 1,
        opinion: "Apa apaan ini?",
        image: "assets/images/testimonial-2.jpg"
    }
]


function allTestimonials(){
    dataHTML = ""
    Data.forEach((item) => {
        dataHTML +=  `<div class="testimonial-item">
        <img src="${item.image}" alt="testimonial-image" class="testimonial-image"/>
        <p class="testimonial-description">${item.opinion}</p>
        <p class="${item.author}">- John</p>
        <p class="testimonial-rating">${item.rating} <i class="fa-solid fa-star" style="color: #000;"></i></p>
    </div>`
    });
    
    document.getElementById("testimonial-list").innerHTML = dataHTML;
}

allTestimonials()

function filteredTestimonials(rating){
    dataHTML = ""

    const filteredData = Data.filter((item) => {
        return rating == item.rating
    })
    if (filteredData.length === 0) {
        dataHTML += `<h1>Data not found!</h1>`;
      } else {
            filteredData.forEach((item) => {
                dataHTML += `<div class="testimonial-item">
                <img src="${item.image}" alt="testimonial-image" class="testimonial-image"/>
                <p class="testimonial-description">${item.opinion}</p>
                <p class="${item.author}">- John</p>
                <p class="testimonial-rating">${item.rating} <i class="fa-solid fa-star" style="color: #000;"></i></p>
            </div>`
            })
        }

    document.getElementById("testimonial-list").innerHTML = dataHTML;
}




