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

class Testimonials {
     constructor(image, opinion, author){
          this.image = image;
          this.opinion = opinion;
          this.author = author;
     }

}

class PersonalTestimonial extends Testimonials{
     constructor(image, opinion, author){
          super(image, opinion, author);
     }

     get review(){
          let result = `<div class="testimonial-item">
                              <img src="${this.image}" alt="testimonial-image" class="testimonial-image"/>
                              <p class="testimonial-description">${this.opinion}</p>
                              <p class="author">- ${this.author}</p>
                         </div>`
          return result
     }
}
class CompanyTestimonial extends Testimonials{
    constructor(image, opinion, author){
         super(image, opinion, author);
        this.author = author + " Company"
    }

    get review(){
         let result = `<div class="testimonial-item">
                             <img src="${this.image}" alt="testimonial-image" class="testimonial-image"/>
                             <p class="testimonial-description">${this.opinion}</p>
                             <p class="author">- ${this.author}</p>
                        </div>`
         return result
    }
}

const testimonial1 = new PersonalTestimonial("assets/images/testimonial-3.jpg", "Lorem ipsum dolor sit amet!", "John");
const testimonial2 = new PersonalTestimonial("assets/images/testimonial-2.jpg", "Lorem ipsum dolor sit amet!", "Sarah");
const testimonial3 = new CompanyTestimonial("assets/images/testimonial-1.jpg", "Lorem ipsum dolor sit amet!", "Robert");

let data = [testimonial1, testimonial2, testimonial3]
let dataHTML = ""
for(let i = 0; i < data.length; i++){
    dataHTML += data[i].review
}
console.log(dataHTML)
document.getElementById("testimonial-list").innerHTML = dataHTML
