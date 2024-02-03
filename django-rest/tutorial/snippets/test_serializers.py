#Before we go any further we'll familiarize ourselves with using our new Serializer class

from snippets.models import Snippet
from snippets.serializers import SnippetSerializer
from rest_framework.renderers import JSONRenderer
from rest_framework.parsers import JSONParser

#let's create a couple of code snippets to work with
snippet = Snippet(code='foo = "bar"\n')
snippet.save()

snippet = Snippet(code='print("hello, world")\n')
snippet.save()

#Let's take a look at serializing one of those instances.
serializer = SnippetSerializer(snippet)
print(serializer.data)
# {'id': 2, 'title': '', 'code': 'print("hello, world")\n', 'linenos': False, 'language': 'python', 'style': 'friendly'}

#To finalize the serialization process we render the data into json.
content = JSONRenderer().render(serializer.data)
print(content)
# b'{"id": 2, "title": "", "code": "print(\\"hello, world\\")\\n", "linenos": false, "language": "python", "style": "friendly"}'


#Deserialization is similar. First we parse a stream into Python native datatypes..
import io

stream = io.BytesIO(content)
data = JSONParser().parse(stream)

# we restore those native datatypes into a fully populated object instance.
serializer = SnippetSerializer(data=data)
print(serializer.is_valid())
# True
print(serializer.validated_data)
# OrderedDict([('title', ''), ('code', 'print("hello, world")\n'), ('linenos', False), ('language', 'python'), ('style', 'friendly')])
print(serializer.save())
# <Snippet: Snippet object>

#We can also serialize querysets instead of model instances. To do so we simply add a many=True flag to the serializer arguments.
serializer = SnippetSerializer(Snippet.objects.all(), many=True)
print(serializer.data)
# [OrderedDict([('id', 1), ('title', ''), ('code', 'foo = "bar"\n'), ('linenos', False), ('language', 'python'), ('style', 'friendly')]), OrderedDict([('id', 2), ('title', ''), ('code', 'print("hello, world")\n'), ('linenos', False), ('language', 'python'), ('style', 'friendly')]), OrderedDict([('id', 3), ('title', ''), ('code', 'print("hello, world")'), ('linenos', False), ('language', 'python'), ('style', 'friendly')])]


# after refactoring the serializer code class using modelserializer
from snippets.serializers import SnippetSerializer
serializer = SnippetSerializer()
print(repr(serializer))