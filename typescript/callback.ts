type Callback = (n1: number, n2: number) => number;

function add(n1: number, n2: number): number {
	return n1 + n2;
}

function multiply(n1: number, n2: number): number {
	return n1 * n2;
}

function exponent(n1: number, n2: number): number {
	return n1 ** n2;
}

function printResult(n1: number, n2: number, callback: Callback) {
	const result = callback(n1, n2);
	console.log(result);
}

printResult(5, 5, add);
printResult(5, 5, multiply);
printResult(5, 5, exponent);
