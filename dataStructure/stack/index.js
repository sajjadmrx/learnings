class Stack {
    constructor() {
        this.items = []
    }

    insert(item) {
        this.items.push(item)
    }

    remove() {
        this.items.pop()
    }


    toArray() {
        return this.items
    }
}

const stack = new Stack()
stack.insert("A")
stack.insert("B")
stack.insert("C")
stack.insert("D")
stack.remove()
stack.remove()
stack.remove()
console.log(stack.toArray())