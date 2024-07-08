# mutn(1) completion

_path_gen() {
    local mutnPath="$HOME/.local/share/libmutton"
    local globStatus=$(shopt -p globstar)
    [ -z "${globStatus##*u*}" ] && shopt -s globstar  # if recursive globbing is disabled, enable it
    local fullPaths=("$mutnPath"/**/*)
    $globStatus  # set recursive globbing to user default
    for scanPath in "${fullPaths[@]}"; do
      # exclude directories
      [ -f "$scanPath" ] && trimmedPaths+=("${scanPath#$mutnPath}")
    done
    if [ "${trimmedPaths[0]}" == '' ]; then trimmedPaths[0]=help; fi
} &&

_mutnCompletions() {
  local cur=${COMP_WORDS[COMP_CWORD]}
  local prev=${COMP_WORDS[COMP_CWORD-1]}

  case $prev in
    mutn )
      while read -r; do ITEM=${REPLY// /\\ }; COMPREPLY+=( "$ITEM" ); done < <( compgen -W "$(printf "'%s' " "${trimmedPaths[@]}")" -- "$cur" )
      ;;
    /* )
      while read -r; do COMPREPLY+=( "$REPLY" ); done < <( compgen -W "copy edit gen add shear" -- "$cur" )
      ;;
    add )
      while read -r; do COMPREPLY+=( "$REPLY" ); done < <( compgen -W "password note folder" -- "$cur" )
      ;;
    copy )
      while read -r; do COMPREPLY+=( "$REPLY" ); done < <( compgen -W "password username totp url note" -- "$cur" )
      ;;
    edit )
      while read -r; do COMPREPLY+=( "$REPLY" ); done < <( compgen -W "password username totp url note rename" -- "$cur" )
      ;;
    gen )
      while read -r; do COMPREPLY+=( "$REPLY" ); done < <( compgen -W "update" -- "$cur" )
      ;;
  esac

} &&
_path_gen && complete -F _mutnCompletions mutn
