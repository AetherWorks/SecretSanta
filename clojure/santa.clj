(def people ["Ally" "Danil" "Greg" "Gus" "Isabel" "Peter" "Renee" "Rob" "Roger" "Shannon"])

(def numPeople (count people))

(def indices (into [] (take numPeople (range))))

(dotimes [i numPeople]
  (print (str (get people i) " -> "))

  (def n (.indexOf indices i))

  (def otherIndices (concat (take n indices) (subvec indices (+ n 1))))

  (def gifteeIndex (nth otherIndices (rand-int (count otherIndices))))

  (println (get people gifteeIndex))

  (def indexToRemove (.indexOf indices gifteeIndex))

  (def indices (vec (concat (subvec indices 0 indexToRemove) (subvec indices (+ 1 indexToRemove)))))

  )
