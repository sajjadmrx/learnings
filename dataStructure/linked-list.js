class Node {
    constructor(data) {
        this.data = data;
        this.next = null
    }
}

class LinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
        this.size = 0;
    }


    prepend(data) {
        let n = new Node(data)
        if (this.size === 0) {
            this.head = n;
            this.tail = n
        } else {
            n.next = this.head
            this.head = n
        }
        this.size++
    }


    append(data) {
        let n = new Node(data)
        if (this.size === 0) {
            this.head = n;
            this.tail = n;
        } else {
            const temp = this.tail;
            this.tail = n;
            temp.next = this.tail
        }
        this.size++
    }
}

const ll = new LinkedList()
ll.append("A")
ll.append("B")
ll.append("C")
console.log(ll.head)