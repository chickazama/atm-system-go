import Navbar from "./components/Navbar.js";

customElements.define("navbar-component", Navbar);

window.addEventListener("load", async () => {
    const main = document.getElementById("main");
    main.appendChild(new Navbar());
})