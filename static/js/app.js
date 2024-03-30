window.addEventListener("load", async () => {
    // const main = document.getElementById("main");
    const body = {
        emailAddress: "test@example.com",
        password: "password123",
        firstName: "test",
        lastName: "user",
        username: "testUser"
    };
    const res = await fetch("/api/users", {
        method: "POST",
        headers: {
            "Content-Type": "application/json" 
        },
        body: JSON.stringify(body)
    })
    const data = await res.json()
    console.log(data);
})