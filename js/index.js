let humbergerIsOpen = false;
console.log("bisa")

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