export async function CreateUserDummyAsync() {
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
}

export async function CreateAccountDummyAsync() {
    const body = {
        accountNumber: "12345678",
        balance: 100
    };
    const res = await fetch("/api/accounts", {
        method: "POST",
        headers: {
            "Content-Type": "application/json" 
        },
        body: JSON.stringify(body)
    })
    const data = await res.json()
    console.log(data);
}
