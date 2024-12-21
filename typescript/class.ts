class Point {
	private x: number;
	private y: number;

	constructor(x?: number, y?:number) {
		this.x = x;
		this.y = y;
	}

	getX() {
		return this.x;
	}

	getY() {
		return this.y;
	}

	setX(x: number) {
		this.x = x;
	}

	setY(y: number) {
		this.y = y;
	}
}

let point = new Point(0, 0);
point.setX(10);
point.setY(20);

console.log(point.getX());
console.log(point.getY());
