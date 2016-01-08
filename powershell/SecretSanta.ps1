# SecretSanta.ps1
# Script that accepts list of names and then determines a secret
# santa match for each person

Param(
    [Parameter(Mandatory=$true)][string[]]$names
)

# Create function that returns $true if for every value in $listA,
# the corresponding value at the same position in $listB is not the 
# same. Returns $false otherwise
Function Compare-Lists ( [string[]]$listA, [string[]]$listB )
{
    $i = 0
    ForEach ($item in $listA) {
        If ($listB[$i] -eq $item) {
            $false
            Return
        }
        $i++
    }
    $true
}

$result = $false
While (!$result) {
    # Shuffle $names list to get $receivers list
    $receivers = $names | Sort-Object {Get-Random}

    # Compare $names and $receivers to see if they create a 
    # valid secret santa output. If they do not, loop will
    # repeat
    $result = Compare-Lists -listA $names -listB $receivers
}

# Print results
Write-Host "Secret Santa Pairings:" -ForegroundColor Red
$i = 0
ForEach ($name in $names) {
    Write-Host ("$name buys a gift for " + $receivers[$i]) -ForegroundColor Green
    $i++
}
Write-Host "~~~~Happy Holidays~~~~" -ForegroundColor Red