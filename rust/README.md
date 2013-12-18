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
      * Is size of Nx equal to 1
         * Re-run - no solution can be found (we're stuck with one person buying for themselves)
      * Shuffle N'
      * Loop
  * Add pair of names to picked pairs (P)
* Return P
