#include <list>


struct TNode {
    TNode(char name, size_t numOfParents, TNode* parent, TNode* child) {
        this->name = name;
        this->numOfParents = numOfParents;

        if (parent != nullptr) {
            parents.push_back(parent);
        }
        if (child != nullptr) {
            children.push_back(child);
        }
    }

    char name;
    size_t numOfParents;
    std::list<TNode*> parents;
    std::list<TNode*> children;
};


class dag{
public:
    dag();
    virtual ~dag();


    void InsertNode(char currNode, char nextNode);
    bool Contains(char name);
    TNode* FindNode(char name);



private:
    TNode* findNodeImpl(char name, TNode* node);
    void addSuccessor(TNode* node, char name);
    TNode* MakeNode(char name, size_t numOfPars, TNode* parent, TNode* child);

    TNode* start;
    TNode* end;
};