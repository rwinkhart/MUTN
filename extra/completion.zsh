#compdef mutn

mutnPath="$HOME/.local/share/libmutton"
fullPaths=("$mutnPath"/**/*)
trimmedPaths=()
for scanPath in "${fullPaths[@]}"; do
  # exclude directories
  [ -f "$scanPath" ] && trimmedPaths+=("${scanPath#$mutnPath}")
done
[[ -z $trimmedPaths ]] && trimmedPaths=(help)

case ${words[-2]} in
  mutn )
    compadd $trimmedPaths
    ;;
  /* )
    compadd {copy,edit,gen,add,shear}
    ;;
  add )
    compadd {password,note,folder}
    ;;
  copy )
    compadd {password,username,url,note}
    ;;
  edit )
    compadd {password,username,url,note,rename}
    ;;
  gen )
    compadd update
    ;;
esac
