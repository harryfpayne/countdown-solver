# Countdown solver

This is a solver for the game show countdown. I've made a cli and an api, I also might make a frontend in the future.

I implemented my own itertools because it felt like an interesting challenge.

The letters round is trivial, just permutate all the letters, doing a binary search into a dictionary for each.

The numbers round is more difficult, it's the same permutation but also all combinations of operations between those 
numbers, and then to be complete I also search through all possible bracket positions incase the order of operations 
matters in reaching the answer. The numbers round won't go through all permutations in under 30 seconds but will almost 
always find an answer, so I'm returning them on a channel.