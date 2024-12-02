import 'dart:convert';
import 'dart:io';


Future<void> main () async {
  var file = File('inputs/input.txt');
  Stream<String> lines = file.openRead()
    .transform(utf8.decoder)
    .transform(LineSplitter());

  var last_seen = 0; 
  var counter = -1;
  int current;
  try{
    await for (var line in lines){
      current = int.parse(line);
      if (current > last_seen){
        counter++;
        print('${current} (increased) ');
      } else{
        
        print('${current} (decreased) ');
      }
      last_seen = current;
    }
  } catch (e){
    
    print('Error: $e');
  }

  print('\n final counter: $counter'); 
  
}