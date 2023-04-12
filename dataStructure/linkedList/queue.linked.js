class QueueLinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
        this.length = 0;
    }


    enqueue(data) {
        const newElement = {
            data,
            next: null
        }
        if (!this.head) {
            this.head = newElement;
            this.tail = newElement
        } else {
            this.tail.next = newElement;
            this.tail = newElement
        }
    }

    toArray() {
        return this.head
    }

}

const ql = new QueueLinkedList()
ql.enqueue("A")
ql.enqueue("B")
ql.enqueue("C")
ql.enqueue("D")
console.log(ql.toArray())