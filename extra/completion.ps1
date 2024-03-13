Class libmuttonEntries : System.Management.Automation.IValidateSetValuesGenerator {
    [string[]] GetValidValues() {
        $entryPath = (Resolve-Path '~/AppData/Local/libmutton/entries').Path
        $entryNames = If (Test-Path $entryPath) {(Get-ChildItem -Path $entryPath -Recurse).FullName.Substring($entryPath.Length) -replace '\\', '/'}
        return [string[]] $entryNames
    }
}

function MUTNArgumentCompleter {
    param ( $commandName,
            $parameterName,
            $wordToComplete,
            $commandAst,
            $fakeBoundParameters )

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
        [Parameter(Mandatory=$false, Position = 0)]
        [ValidateSet([libmuttonEntries])]
        [string]$entry,

        [Parameter(Position = 1)]
        [ArgumentCompletions('add', 'gen', 'edit', 'copy', 'shear')]
        $argument,

        [Parameter(Position = 2)]
        [ArgumentCompleter({ MUTNArgumentCompleter @args })]
        $option
      )
}
