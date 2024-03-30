window.addEventListener("load", async () => {
    // const main = document.getElementById("main");
    const body = {
        emailAddress: "test2@example.com",
        password: "password123",
        firstName: "test",
        lastName: "user2",
        username: "testUser2"
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
    // const body = {
    //     accountNumber: "12345678",
    //     balance: 100
    // };
    // const res = await fetch("/api/accounts", {
    //     method: "POST",
    //     headers: {
    //         "Content-Type": "application/json" 
    //     },
    //     body: JSON.stringify(body)
    // })
    // const data = await res.json()
    // console.log(data);
})