const array1 = [1,2,3,4];
const array2 = [5,6,7,8];

array1.push(...array2);

const person = {
	name: 'Dom',
	age: 44
};

const personCopy = {...person};

console.log(array1);
console.log(personCopy);
