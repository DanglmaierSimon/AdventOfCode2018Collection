#include <cassert>

#include "dag.hpp"

dag::dag() {
    start = nullptr;
    end = nullptr;
}





void dag::addSuccessor(TNode* node, char name){
    assert(node != nullptr);

    TNode* newNode = MakeNode(name, 1, node, nullptr);

    node->children.push_back(newNode);
}

TNode* dag::MakeNode(char name, size_t numOfPars, TNode* parent, TNode* child){
    TNode* newNode = new TNode{name, numOfPars, parent, child};

    assert(newNode != nullptr);

    return newNode;
}

TNode* dag::FindNode(char name){
    return findNodeImpl(name, start);
}

TNode* findNodeImpl(char name, TNode* node) {
    if (node == nullptr || node->name != name) {
        return nullptr;
    }
    else if (node->name == name) {
        return node;
    }
    else {
        for (auto const & e : node->children) {
            if (findNodeImpl(name,e) != nullptr){
                return e;
            }
        }
    }

    return nullptr;

}

void dag::InsertNode(char currNode, char nextNode) {

}