%% @author Angus Macdonald (amacdonald AT aetherworks.com)
%% @doc Utility for picking secret santa selections from a group of people.
%% @version 1.0

-module(santa).

-export([select/1]).

-spec(select(list()) -> list()).
%% @doc Takes a list of peoples names, and makes secret assignments
%% to determine who each person's secret santa is.
select(People) ->
	print_list("Unshuffled: ", People),

	ShuffledPeople = shuffle(People),
	print_list("Shuffled: ", ShuffledPeople),
	
	RotatedPeople = rotate(ShuffledPeople),

	print_list("Rotated: ", RotatedPeople),

	Assignments = assign(ShuffledPeople, RotatedPeople),

	Assignments.


%% @doc Takes two equal size lists and zips them, to create the
%% assignments for each persons secret santa.
-spec(assign(list(), list()) -> list()).
assign(ListA, ListB) ->
	lists:zip(ListA, ListB).

%% @doc Takes a list and returns a new list, which is left rotated by 1.
-spec(rotate(list()) -> list()).
rotate([]) ->
	[];
rotate(List)->
	{Left, Right} = lists:split(1, List), 
	Right ++ Left.
 

%% @doc Utility to print the contents of a list, prepended by a prefix.
 print_list(Prefix, List) ->
 	io:format(Prefix),
	ListPrinter = fun(E) -> io:format("~p, ",[E]) end,
 	lists:foreach(ListPrinter,List),
 	io:format("~n").

%% Utility to shuffle a list and return the new list.
shuffle(List) -> shuffle(List, []).
shuffle([], Acc) -> Acc;
shuffle(List, Acc) ->
  {Leading, [H | T]} = lists:split(random:uniform(length(List)) - 1, List),
  shuffle(Leading ++ T, [H | Acc]).