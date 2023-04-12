class StackLinked {
    constructor() {
        this.top = null
        this.length = 0
    }

    push(data) {
        const newElement = {
            data,
            next: null
        }
        if (!this.top) {
            this.top = newElement
        } else {
            newElement.next = this.top;
            this.top = newElement
        }


        this.length++;
    }

    pop() {
        if (!this.top) return null;
        const data = this.top.data;
        this.top = this.top.next;
        this.length--;
        return data
    }

    toArray() {
        return this.top
    }
}

const sl = new StackLinked()
sl.push("A")
sl.push("B")
sl.push("C")
sl.push("D")
sl.pop()
sl.pop()
console.log(sl.toArray())



