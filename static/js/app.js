window.addEventListener("load", async () => {
    const main = document.getElementById("main");
    const a = document.createElement("a");
    a.href = "/auth/github";
    a.innerText = "Sign in with GitHub";
    main.appendChild(a);
})