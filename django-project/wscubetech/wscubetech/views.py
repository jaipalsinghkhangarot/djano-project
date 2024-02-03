from django.http import HttpResponse,HttpResponseRedirect
from django.shortcuts import render,redirect

def homePage(request):
    return render(request,"index.html")

def userFormGet(request):
    sum = 0
    try:
        n1 = int(request.GET['num1'])
        n2 = int(request.GET['num2'])
        sum = n1+n2
    except:
        pass
    return render(request,"userformget.html",{'output':sum})

def submitform(request):
    sum = 0
    data = {}
    try:
        n1 = int(request.POST['num1'])
        n2 = int(request.POST['num2'])
        sum = n1+n2
        data = {'n1':n1,'n2':n2,'output':sum}
    except:
        pass
    return HttpResponse(sum)

def userFormPost(request):
    sum = 0
    data = {}
    try:
        n1 = int(request.POST['num1'])
        n2 = int(request.POST['num2'])
        sum = n1+n2
        data = {'n1':n1,'n2':n2,'output':sum}
    except:
        pass
    return render(request,"userformpost.html",data)

def basicPage(request):
    data = {
        "title":"Home Page",
        "body":"WELCOME TO JAI'S DJANGO PROJECT!!!",
        "course":['PYTHON','JAVA','DJANGO'],
        "numbers":[10,20,30,40,50],
        "student":[{'name':'jaipal','phone':7340300100},{'name':'param','phone':9314638839}]
    }
    return render(request,"basic.html",data)

def aboutUs(request):
    return HttpResponse("Welcome to Hell!!!")

def course(request):
    return HttpResponse("Welcome to dynamic Hell!!!")

def courseDetails(request,courseid):
    return HttpResponse(courseid)