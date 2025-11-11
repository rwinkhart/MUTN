$global:cliMUTN_entriesToSpaceIndicesMap = @{}

if (!$IsWindows) {
    # since the executable has the same name as the completion function (on non-Windows platforms), its path must be stored before the function is registered
    $MUTNExecutablePath = (Get-Command mutn).Path
    $spaceIndicesFilePath = '~/.config/libmutton/psCompletionsCache.json'
} else {
    $spaceIndicesFilePath = '~\AppData\Local\libmutton\psCompletionsCache.json'
}

function cliMUTN-loadSpaceIndices {
    if (Test-Path $global:spaceIndicesFilePath) {
        $jsonContent = Get-Content -Path $global:spaceIndicesFilePath -Raw
        $global:cliMUTN_entriesToSpaceIndicesMap = ConvertFrom-Json -InputObject $jsonContent -AsHashtable
    }
}

function cliMUTN-saveSpaceIndices {
    $jsonContent = $global:cliMUTN_entriesToSpaceIndicesMap | ConvertTo-Json -Compress
    $jsonContent | Set-Content -Path $global:spaceIndicesFilePath
}

function cliMUTN-getSpaceIndices {
    param ([string]$inputString)
    $indices = @()

    # loop through each character in the string
    for ($i = 0; $i -lt $inputString.Length; $i++) {
        if ($inputString[$i] -eq ' ') {
            # if the character is a space, add the index to the array
            $indices += $i
        }
    }

    return $indices
}

function cliMUTN-getEscapedEntryName {
    param ([int[]]$Indices, [string]$InputString)
    $charArray = $InputString.ToCharArray()

    # loop through each index in the provided array,
    # replacing the character at that index with a space
    foreach ($index in $Indices) {
        $charArray[$index] = ' '
    }

    # convert the character array back to a string
    $outputString = -join $charArray

    # return the modified string with spaces escaped
    return $outputString -replace ' ', '` '
}

function cliMUTN-entryCompleter {
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameter)
    if ($IsWindows) {
        $entryRoot = (Resolve-Path '~\AppData\Local\libmutton\entries').Path
    } else {
        $entryRoot = (Resolve-Path '~/.local/share/libmutton').Path
    }
    try {
        $trimmedPaths = If (Test-Path $entryRoot) {
            (Get-ChildItem -Path $entryRoot -Recurse -File).FullName.Substring($entryRoot.Length) -replace '\\', '/'
        }
    } catch {
        $trimmedPaths = $null # if any errors occur (especially, "You cannot call a method on a null-valued expression", set $trimmedPaths to $null
    }
    if ($null -eq $trimmedPaths) { # if no entries are found, add 'help' to $trimmedPaths
        $trimmedPaths = 'help'
    } else {
        # replace spaces with underscores, tracking the indices of the spaces in a global variable for later restoration of spaces
        $trimmedPaths = $trimmedPaths | ForEach-Object {
            $spaceIndices = cliMUTN-getSpaceIndices -inputString $_
            $replacedEntry = $_ -replace ' ', '_'
            $replacedEntry
            $global:cliMUTN_entriesToSpaceIndicesMap[$replacedEntry] = $spaceIndices
        }
        cliMUTN-saveSpaceIndices
    }
    $trimmedPaths | Where-Object { $_ -like "$wordToComplete*" }
}

function cliMUTN-optionCompleter {
    param ($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameters)

    $possibleValues = @{
        add = @('password', 'note', 'folder')
        copy = @('password', 'username', 'totp', 'url', 'note', 'menu')
        edit = @('password', 'username', 'totp', 'url', 'note', 'rename')
        gen = @('update')
    }

    if ($fakeBoundParameters.ContainsKey('argument')) {
        $possibleValues[$fakeBoundParameters.argument] | Where-Object {
            $_ -like "$wordToComplete*"
        }
    } else {
        $possibleValues.Values | ForEach-Object {$_}
    }
}

function mutn {
    [CmdletBinding()]
    param (
        [Parameter(Position = 0)]
        [ArgumentCompleter({ cliMUTN-entryCompleter @args })]
        [string]$entry,

        [Parameter(Position = 1)]
        [ArgumentCompletions('copy', 'edit', 'gen', 'add', 'shear')]
        [string]$argument,

        [Parameter(Position = 2, ValueFromRemainingArguments=$true)]
        [ArgumentCompleter({ cliMUTN-optionCompleter @args })]
        [string]$option
      )

    # replace placeholder underscores with escaped spaces
    $escapedEntry = cliMUTN-getEscapedEntryName -Indices $global:cliMUTN_entriesToSpaceIndicesMap[$entry] -InputString $entry

    if ($IsWindows) {
        Invoke-Expression -Command ('mutn.exe ' + $escapedEntry, $argument, $option).Trim()
    } else {
        Invoke-Expression -Command ($global:MUTNExecutablePath + ' ' + $escapedEntry, $argument, $option).Trim()
    }
}

cliMUTN-loadSpaceIndices
