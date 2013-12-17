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
	let mut unshuffled_names = names.clone();
	let mut shuffled_names = rng.shuffle(names.clone());

	println!("Picking pairings from: {:?}", unshuffled_names);

	let mut pairs = ~[];

	while shuffled_names.len() > 0 {
		let chosen = unshuffled_names.pop();
		let partner = shuffled_names.pop();

		if( chosen == partner ) {
			unshuffled_names.push(chosen);
			shuffled_names.push(partner);

			if(shuffled_names.len() == 1) {
				break;
			}
			rng.shuffle_mut(shuffled_names);
			loop; 
		}

		pairs.push((chosen,partner));
	}

	if(shuffled_names.len() == 1) {
		println("Restarting - no good solution.");
		secret_santa(names);
	}

	return pairs.clone();
} 

fn print_pairings(pairings: &~[(&str, &str)]) {
	for &pair in pairings.iter() {
			println!("{:?} --- is buying for ---> {:?}", pair.first(), pair.second());
	}
}