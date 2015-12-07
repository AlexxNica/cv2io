/* add more form elements */
var n = 0;

function add(fieldsetName) {
  addedMore = document.getElementById(fieldsetName);
  addedCopy = addedMore.cloneNode(true);

  n++;
  addedCopy.id = fieldsetName + n;

  addedParent = addedMore.parentNode;
  addedParent.insertBefore(addedCopy, addedMore);
}


/*
       var InputType = clone.getElementsByTagName("input");

       for (var i=0; i<InputType.length; i++){
        if( InputType[i].type=='checkbox'){
           InputType[i].checked = false;
       }else{
          InputType[i].value='';
           }
       }
       table.appendChild(clone);

*/
