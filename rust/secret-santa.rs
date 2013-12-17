use std::rand;
use std::rand::Rng;

fn main() {
	println("Secret Santa");

	let names = ~["Angus", "Greg", "Lewis", "Mike", "Isabel", "Rob", "Shannon", "Allan"];

	let pairings = secret_santa(&names);	

	print_pairings(&pairings);
}

fn secret_santa<'r>(names: &~[&'r str]) -> ~[(&'r str,&'r str)] {
	let mut rng = rand::rng();
	let mut shuffled_names = rng.shuffle(names.clone());

	println!("Picking pairings from: {:?}", names);

	let mut pairs = ~[];

	let first = shuffled_names.last().clone();
	shuffled_names.unshift(first);

	while shuffled_names.len() > 1 {
		let chosen = shuffled_names.pop();
		let partner = shuffled_names.last().clone();

		pairs.push((chosen,partner));
	}

	return pairs.clone();
} 

fn print_pairings(pairings: &~[(&str, &str)]) {
	for &pair in pairings.iter() {
			println!("{:?} --- is buying for ---> {:?}", pair.first(), pair.second());
	}
}