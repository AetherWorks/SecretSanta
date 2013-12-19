%% @author Angus Macdonald (amacdonald AT aetherworks.com)
%% @doc Utility for picking secret santa selections from a group of people.

-module(santa).

-export([select/1]).

-spec(select(list()) -> list()).
%% @doc Takes a list of peoples names, shuffles to create a random ordering,
%% then rotates a clone of this list to determine who has who in secret santa.
select(People) ->
	ShuffledPeople = shuffle(People),
	RotatedPeople = rotate(ShuffledPeople),
	lists:zip(ShuffledPeople, RotatedPeople).

%% @doc Takes a list and returns a new list, which is left rotated by 1.
-spec(rotate(list()) -> list()).
rotate([]) ->
	[];
rotate(List)->
	{Left, Right} = lists:split(1, List), 
	Right ++ Left.
 
%% @doc Utility to shuffle a list and return the new list. Accumulates the new list by
%% picking a random element from the existing list, then recursively doing this for each
%% successively smaller list.
%% From: http://chimera.labs.oreilly.com/books/1234000000726/ch07.html#CH07-ET05
shuffle(List) -> shuffle(List, []).
shuffle([], Acc) -> Acc;
shuffle(List, Acc) ->
  {Leading, [H | T]} = lists:split(random:uniform(length(List)) - 1, List),
  shuffle(Leading ++ T, [H | Acc]).
