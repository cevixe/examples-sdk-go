#!/usr/bin/env bash

IFS=$'\n'

mergeUtility="graphql-schema-utilities"

mergedSchema="./.aws-sam/schema/merged.graphql"
tempSchema="./.aws-sam/schema/temp.graphql"
finalSchema="./.aws-sam/schema/final.graphql"

schemasGlobSuffix='*.graphql'

graphqlModules=( $(ls -dl modules/api/schemas/* | awk '{print $9}') )

echo "* Modules: ${#graphqlModules[@]}"
echo "* Globs:"
schemaGlobPattern=""
for module in "${graphqlModules[@]}"
do
  newGlobPattern=$(echo "./$module/$schemasGlobSuffix")
  schemaGlobPattern+="$newGlobPattern,"
  echo "  - $newGlobPattern"
done
schemaGlobPattern=$(echo $schemaGlobPattern | rev | cut -c2- | rev)
schemaGlobPattern="{$schemaGlobPattern}"

eval "$mergeUtility -d -s \"$schemaGlobPattern\" -o $mergedSchema"
eval "sed -E 's/^([[:space:]]*)\"\"\"(.*)\"\"\"/\1# \2/g' $mergedSchema > $tempSchema"

isComment=false
echo "# Auto-generated GraphQL Schema" > $finalSchema
while IFS= read -r line
do
  if [[ "$line" =~ ^([[:space:]]*)\"\"\"\.* ]]; then
    if [[ "$isComment" = true ]]; then
      isComment=false; continue;
    elif [[ "$isComment" = false ]]; then
      isComment=true; continue;
    fi
  fi

  if [[ "$isComment" = true ]]; then
    echo "$line" | sed -E 's/^([[:space:]]*)(.*)/\1# \2/g' >> $finalSchema
  else
    echo "$line" >> $finalSchema
  fi
done < "$tempSchema"

rm $mergedSchema
rm $tempSchema

echo "* GraphQL Schema Generated"
