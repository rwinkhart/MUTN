function cliMUTNEntryCompleter {
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameter)
    #$mutnPath = (Resolve-Path '~/.local/share/libmutton').Path # UNIX testing
    $mutnPath = (Resolve-Path '~/AppData/Local/libmutton/entries').Path
    try {
        $trimmedPaths = If (Test-Path $mutnPath) {
            (Get-ChildItem -Path $mutnPath -Recurse -File).FullName.Substring($mutnPath.Length) -replace '\\', '/' -replace ' ', '` '
        }
    } catch {
        $trimmedPaths = $null # if any errors occur (especially, "You cannot call a method on a null-valued expression", set $trimmedPaths to $null
    }
    if ($null -eq $trimmedPaths) { # if no entries are found, add 'help' to $trimmedPaths
        $trimmedPaths = 'help'
    }
    $trimmedPaths | Where-Object { $_ -like "$wordToComplete*" }
}

function cliMUTNOptionCompleter {
    param ($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameters)

    $possibleValues = @{
        add = @('password', 'note', 'folder')
        copy = @('password', 'username', 'totp', 'url', 'note')
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
        [ArgumentCompleter({ cliMUTNEntryCompleter @args })]
        [string]$entry,

        [Parameter(Position = 1)]
        [ArgumentCompletions('copy', 'edit', 'gen', 'add', 'shear')]
        [string]$argument,

        [Parameter(Position = 2, ValueFromRemainingArguments=$true)]
        [ArgumentCompleter({ cliMUTNOptionCompleter @args })]
        [string]$option
      )

    #Invoke-Expression -Command ('/usr/local/bin/mutn ' + ($entry -replace ' ', '` '), $argument, $option).Trim() # UNIX testing
    Invoke-Expression -Command ('./mutn.exe ' + ($entry -replace ' ', '` '), $argument, $option).Trim()
}
