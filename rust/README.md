To run
---

rust run secret-santa.rs

Algorithm
---

* Create stack of names (N)
* Create stack of shuffled names (N')
* While N is not empty
 * Pop values off N (Nx) and N' (Nx')
 * Compare Nx and Nx'
   * If same
   * Push values Nx and Nx' onto N and N'
   * Shuffle N'
   * Loop 
