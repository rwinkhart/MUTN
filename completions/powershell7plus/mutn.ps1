if (!$IsWindows) {
    # since the executable has the same name as the completion function (on non-Windows platforms), its path must be stored before the function is registered
    $MUTNExecutablePath = (Get-Command mutn).Path
}

function cliMUTNEntryCompleter {
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameter)
    if ($IsWindows) {
        $entryRoot = (Resolve-Path '~\AppData\Local\libmutton\entries').Path
    } else {
        $entryRoot = (Resolve-Path '~/.local/share/libmutton').Path
    }
    try {
        $trimmedPaths = If (Test-Path $entryRoot) {
            (Get-ChildItem -Path $entryRoot -Recurse -File).FullName.Substring($entryRoot.Length) -replace '\\', '/' -replace ' ', [char]0x259d
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

    if ($IsWindows) {
        Invoke-Expression -Command ('mutn.exe ' + ($entry -replace ' ', '` ' -replace [char]0x259d, '` '), $argument, $option).Trim()
    } else {
        Invoke-Expression -Command ($MUTNExecutablePath + ' ' + ($entry -replace ' ', '` ' -replace [char]0x259d, '` '), $argument, $option).Trim()
    }
}
