import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Random;

/**
 * Randomly generates Secret Santa assignments for a given group.
 * <p>
 * All valid possible assignments are equally likely to be generated (uniform distribution).   
 * 
 * @author Michael Zaccardo (mzaccardo@aetherworks.com)
 */
public class SecretSanta {

	private static final Random random = new Random();

	private static final String[] DEFAULT_NAMES = { "Rob", "Ally", "Angus", "Mike", "Shannon", "Greg", "Lewis", "Isabel" };

	public static void main(final String[] args) {
		final String[] names = getNamesToUse(args);

		final List<Integer> assignments = generateAssignments(names.length);

		printAssignmentsWithNames(assignments, names);
	}

	private static String[] getNamesToUse(final String[] args) {
		if (args.length >= 2) {
			return args;
		} else {
			System.out.println("Two or more names were not provided -- using default names.\n");
			return DEFAULT_NAMES;
		}
	}

	private static List<Integer> generateAssignments(final int size) {
		final List<Integer> assignments = generateShuffledList(size);

		while (!areValidAssignments(assignments)) {
			Collections.shuffle(assignments, random);
		}

		return assignments;
	}

	private static List<Integer> generateShuffledList(final int size) {
		final List<Integer> list = new ArrayList<>(size);

		for (int i = 0; i < size; i++) {
			list.add(i);
		}
		
		Collections.shuffle(list, random);

		return list;
	}

	private static boolean areValidAssignments(final List<Integer> assignments) {
		for (int i = 0; i < assignments.size(); i++) {
			if (i == assignments.get(i)) {
				return false;
			}
		}

		return true;
	}

	private static void printAssignmentsWithNames(final List<Integer> assignments, final String[] names) {
		for (int i = 0; i < assignments.size(); i++) {
			System.out.println(names[i] + " --> " + names[assignments.get(i)]);
		}
	}
}