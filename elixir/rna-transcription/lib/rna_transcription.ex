defmodule RnaTranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    complements = %{?G => ?C, ?C => ?G, ?T => ?A, ?A => ?U}
    Enum.map(dna, fn x -> Map.get(complements, x) end)
  end
end
