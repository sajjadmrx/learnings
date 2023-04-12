class Queue {
    constructor() {
        this.items = []
    }

    enQueue(data) {
        this.items.unshift(data)
    }

    deQueue() {
        return this.items.pop()
    }

    toArray() {
        return this.items
    }


}

const x = new Queue()

x.enQueue("A") //first out
x.enQueue("B")
x.enQueue("C")
x.enQueue("D")
x.deQueue()
x.deQueue()
x.enQueue("E")
x.deQueue()
console.log(x.toArray())