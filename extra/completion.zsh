#compdef mutn

mutn_path="$HOME/.local/share/libmutton"
full_paths=("$mutn_path"/**/*)
trimmed_paths=()
for scan_path in "${full_paths[@]}"; do
  trimmed_paths+=("${scan_path#$mutn_path}")
done
[[ -z $trimmed_paths ]] && trimmed_paths=(help)

case ${words[-2]} in
  mutn )
    compadd $trimmed_paths
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
