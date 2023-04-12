class User {

    constructor(
        public email: string,
        public next: User | null,
        public back: User | null
    ) {
    }

}


class LinkedListUser {
    private head: User | null = null
    private tail: User | null = null
    constructor() {

    }


    append(email: string) {
        const user = new User(email, null, this.head)
        if (!this.head) {
            this.head = user
            this.tail = this.head //{ ...user }
        } else {
            if (this.tail) {
                this.tail.next = user
                user.back = this.tail
                this.tail = user
            }
        }

    }

    toArray() {
        return this.head
    }
}

const users = new LinkedListUser()
users.append("fake@gmail.com")
users.append("fake2@gmail.com")
users.append("fake3@gmail.com")

console.log(users.toArray()?.next?.next?.back)