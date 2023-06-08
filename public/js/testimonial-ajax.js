const getDataFromServer = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://api.npoint.io/27d19e543b1fedda6ee0", true);

  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Cannot loading data from server!");
    }
  };
  xhr.onerror = () => {
    reject("Internet disconnected!");
  };
  xhr.send();
});

async function getAllTestimonials() {
  try {
    const response = await getDataFromServer;
    console.log(response);

    let dataHTML = "";
    response.forEach((item) => {
      dataHTML += `<div class="testimonial-item">
            <img src="${item.image}" alt="testimonial-image" class="testimonial-image"/>
            <p class="testimonial-description">${item.opinion}</p>
            <p class="${item.author} text-end">- John</p>
            <p class="testimonial-rating">${item.rating} <i class="fa-solid fa-star text-dark"></i></p>
        </div>`;
    });

    document.getElementById("testimonial-list").innerHTML = dataHTML;
  } catch (err) {
    console.log("Error: " + err);
  }
}

async function getFilteredTestimonials(rating) {
  try {
    const response = await getDataFromServer;
    console.log(response);

    let dataHTML = "";

    const filteredData = response.filter((item) => {
      return item.rating === rating;
    });

    if (filteredData.length === 0) {
      dataHTML += `<h1>Data not found!</h1>`;
    } else {
      filteredData.forEach((item) => {
        dataHTML += `<div class="testimonial-item">
                <img src="${item.image}" alt="testimonial-image" class="testimonial-image"/>
                <p class="testimonial-description">${item.opinion}</p>
                <p class="${item.author} text-end">- John</p>
                <p class="testimonial-rating">${item.rating} <i class="fa-solid fa-star text-dark"></i></p>
            </div>`;
      });
    }

    document.getElementById("testimonial-list").innerHTML = dataHTML;
  } catch (err) {
    console.log("Error: " + err);
  }
}

getAllTestimonials();
