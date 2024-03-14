function MUTNEntryCompleter {
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameter)
    $entryPath = (Resolve-Path '~/.local/share/libmutton').Path # UNIX testing
    #$entryPath = (Resolve-Path '~/AppData/Local/libmutton/entries').Path
    $entryNames = If (Test-Path $entryPath) {(Get-ChildItem -Path $entryPath -Recurse).FullName.Substring($entryPath.Length) -replace '\\', '/' -replace ' ', '` '}
    $entryNames | Where-Object { $_ -like "$wordToComplete*" }
}

function MUTNOptionCompleter {
    param ($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameters)

    $possibleValues = @{
        add = @('password', 'note', 'folder')
        copy = @('password', 'username', 'url', 'note')
        edit = @('password', 'username', 'url', 'note', 'rename')
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
        [ArgumentCompleter({ MUTNEntryCompleter @args })]
        [string]$entry,

        [Parameter(Position = 1)]
        [ArgumentCompletions('add', 'gen', 'edit', 'copy', 'shear')]
        [string]$argument,

        [Parameter(Position = 2)]
        [ArgumentCompleter({ MUTNOptionCompleter @args })]
        [string]$option
      )

    Invoke-Expression -Command ('/usr/local/bin/mutn ' + ($entry -replace ' ', '` '), $argument, ($option -replace ':', '-')).Trim() # UNIX testing
    #Invoke-Expression -Command ('mutn.exe ' + ($entry -replace ' ', '` '), $argument, $option).Trim()
}
