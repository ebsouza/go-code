FOLDER=test_data

mkdir -p $FOLDER/inside_1
mkdir -p $FOLDER/inside_2


for f in {0..10}
do
    echo "Some phrase here" > "$FOLDER/file_$f.log"
    echo "Some phrase here" > "$FOLDER/inside_1/file_$f.log"
    echo "Some phrase here" > "$FOLDER/inside_2/file_$f.log"
done

for f in {0..10}
do
    echo "Some phrase here" > "$FOLDER/file_$f.dat"
    echo "Some phrase here" > "$FOLDER/inside_1/file_$f.dat"
    echo "Some phrase here" > "$FOLDER/inside_2/file_$f.dat"
done