use std::rand;
use std::rand::Rng;

fn main() {
	println("Secret Santa");

	let names = ~["Angus", "Greg", "Lewis", "Mike", "Isabel", "Rob", "Shannon", "Allan"];

	secret_santa(names, pairs);	
}

fn secret_santa(names: ~[&str], pairs: ~[(&str, &str)]) {
	let mut rng = rand::rng();
	let mut unshuffled_names = names.clone();
	let mut shuffled_names = rng.shuffle(names.clone());

	println!("Picking pairings from: {:?}", unshuffled_names);

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
		secret_santa(names.clone());
	} else {
		print_pairings(pairs);
	}

	return pairs; 
} 

fn print_pairings(pairings: ~[(&str, &str)]) {
	for &pair in pairings.iter() {
			println!("{:?} --- is buying for ---> {:?}", pair.first(), pair.second());
	}
}