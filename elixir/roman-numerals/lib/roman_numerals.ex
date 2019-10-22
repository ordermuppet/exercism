defmodule RomanNumerals do
  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) do
    thousands_place = div(number, 1000)
    hundreds_place = div(rem(number,1000), 100)
    tens_place = div(rem(number, 100), 10)
    ones_place = rem(number, 10)

    String.duplicate("M", thousands_place)
      <> to_numeral(hundreds_place, "C", "D", "M")
      <> to_numeral(tens_place, "X", "L", "C")
      <> to_numeral(ones_place, "I", "V", "X")
  end

  defp to_numeral(number, one, five, ten) do
    cond do
      number == 0 ->
        ""
      number < 4 ->
        String.duplicate(one, number)
      number == 4 ->
        one <> five
      number == 5 ->
        five
      number < 9 ->
        five <> String.duplicate(one, number - 5)
      true ->
        one <> ten
    end
  end
end
