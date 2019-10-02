defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence 
    |> clean()
    |> reduce_to_count()
  end

  defp clean(sentence) do
    sentence
    |> String.downcase()
    |> String.split(~r/[^[:alnum:]-]/u, trim: true)
  end

  defp reduce_to_count(words) do
    Enum.reduce(words, %{}, fn (x, acc) ->
      Map.update(acc, x, 1, &(&1 + 1))
    end)
  end
end
